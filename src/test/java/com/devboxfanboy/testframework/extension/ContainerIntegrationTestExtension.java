package com.devboxfanboy.testframework.extension;

import java.lang.reflect.Field;
import java.util.ArrayList;
import java.util.List;

import org.apache.commons.collections4.CollectionUtils;
import org.apache.commons.lang3.time.StopWatch;
import org.hibernate.Session;
import org.hibernate.SessionFactory;
import org.junit.jupiter.api.extension.AfterAllCallback;
import org.junit.jupiter.api.extension.BeforeAllCallback;
import org.junit.jupiter.api.extension.BeforeEachCallback;
import org.junit.jupiter.api.extension.ExtensionContext;
import org.mockito.Mock;
import org.mockito.Mockito;
import org.mockito.Spy;

import com.devboxfanboy.FlywayUtil;
import com.devboxfanboy.hibernate.HibernateUtil;
import com.devboxfanboy.testframework.Components;
import com.devboxfanboy.testframework.annotation.Container;
import com.devboxfanboy.testframework.container.impl.IntegrationTestContainer;
import com.devboxfanboy.testframework.container.impl.IntegrationTestWeldContainer;
import jakarta.inject.Inject;

public class ContainerIntegrationTestExtension implements BeforeAllCallback, BeforeEachCallback, AfterAllCallback {

    private IntegrationTestContainer testContainer;
    boolean isInitialized;

    @Override
    public void beforeAll(ExtensionContext extensionContext) throws Exception {
        isInitialized = false;
        StopWatch stopWatch = new StopWatch();
        stopWatch.start();
        FlywayUtil.setUp();
        SessionFactory sessionFactory = HibernateUtil.getSessionFactory();
        Session session = sessionFactory.openSession();
        session.beginTransaction();
        session.createNativeMutationQuery(
                        "insert into opistsschema.CASBIN_RULE (PTYPE,V0,V1,V2,V3,V4,V5,version,ID) values ('p','alice','test','testpermission',NULL,NULL,NULL,0,default)")
                .executeUpdate();
        session.getTransaction().commit();
        session.close();
        stopWatch.stop();
        System.out.println("Database setup in " + stopWatch.getTime() + "ms");
    }

    @Override
    public void afterAll(ExtensionContext extensionContext) {
        if (testContainer != null) {
            testContainer.shutdown();
        }
    }

    @Override
    public void beforeEach(ExtensionContext extensionContext) throws Exception {
        if (isInitialized) {
            return;
        }
        StopWatch stopWatch = new StopWatch();
        stopWatch.start();
        Object testInstance = extensionContext.getRequiredTestInstance();
        Class<?> testClass = testInstance.getClass();

        List<Class<?>> componentClasses = new ArrayList<>();
        for (Field field : testClass.getDeclaredFields()) {
            if (field.isAnnotationPresent(Container.class)) {
                if (!field.getType().equals(Components.class)) {
                    throw new IllegalArgumentException("Field annotated with @Container must be of type Components");
                }
                field.setAccessible(true);
                Components components = (Components) field.get(testInstance);
                componentClasses.addAll(components.getConfiguration().getClasses());
            }
        }
        testContainer = new IntegrationTestContainer(new IntegrationTestWeldContainer(componentClasses));
        testContainer.start();

        List<Class<?>> mockClasses = new ArrayList<>();
        List<Object> mockInstances = new ArrayList<>();
        for (Field field : testClass.getDeclaredFields()) {
            if (field.isAnnotationPresent(Mock.class)) {
                field.setAccessible(true);
                Object mockInstance = Mockito.mock(field.getType());
                field.set(testInstance, mockInstance);
                mockClasses.add(field.getType());
                mockInstances.add(mockInstance);
            }
        }

        List<Class<?>> spyClasses = new ArrayList<>();
        List<Object> spyInstances = new ArrayList<>();
        for (Field field : testClass.getDeclaredFields()) {
            if (field.isAnnotationPresent(Spy.class)) {
                field.setAccessible(true);
                Object instanceOfField = testContainer.get(field.getType());
                Object spyInstance = Mockito.spy(instanceOfField);
                field.set(testInstance, spyInstance);
                spyClasses.add(field.getType());
                spyInstances.add(spyInstance);
            }
        }

        for (Field field : testClass.getDeclaredFields()) {
            if (field.isAnnotationPresent(Inject.class)) {
                field.setAccessible(true);
                Object instanceOfField = testContainer.get(field.getType());
                field.set(testInstance, instanceOfField);
                if (CollectionUtils.isNotEmpty(mockClasses) || CollectionUtils.isNotEmpty(spyClasses)) {
                    Field[] declaredFields = field.getType().getDeclaredFields();
                    for (Field declaredField : declaredFields) {
                        if (mockClasses.contains(declaredField.getType())) {
                            declaredField.setAccessible(true);
                            Object mockInstance = mockInstances.get(mockClasses.indexOf(declaredField.getType()));
                            declaredField.set(instanceOfField, mockInstance);
                        } else if (spyClasses.contains(declaredField.getType())) {
                            declaredField.setAccessible(true);
                            Object spyInstance = spyInstances.get(spyClasses.indexOf(declaredField.getType()));
                            declaredField.set(instanceOfField, spyInstance);
                        }
                    }
                }
            }
        }
        stopWatch.stop();
        System.out.println("Test instance initialized in " + stopWatch.getTime() + "ms");
    }
}
