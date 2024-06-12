package com.devboxfanboy.testframework.container.impl;

import java.util.List;

import org.apache.commons.lang3.time.StopWatch;
import org.jboss.weld.environment.se.Weld;
import org.jboss.weld.environment.se.WeldContainer;
import org.jboss.weld.proxy.WeldClientProxy;

import com.devboxfanboy.testframework.container.TestContainer;

public class IntegrationTestWeldContainer implements TestContainer {

    private WeldContainer container;
    private final List<Class<?>> components;

    boolean started = false;

    public IntegrationTestWeldContainer(List<Class<?>> components) {
        this.components = components;
    }

    @Override
    public Object get(Class<?> clazz) {
        Object instanceOfField = container.select(clazz).get();
        if (instanceOfField instanceof WeldClientProxy) {
            instanceOfField = ((WeldClientProxy) instanceOfField).getMetadata().getContextualInstance();
        }
        return instanceOfField;
    }

    @Override
    public void shutdown() {
        container.shutdown();
    }

    @Override
    public void start() {
        if (started) {
            return;
        }
        StopWatch stopWatch = new StopWatch();
        stopWatch.start();
        Weld weld = new Weld();
        weld.addBeanClasses(components.toArray(new Class<?>[0]));
        container = weld.initialize();
        for (Class<?> clazz : components) {
            container.select(clazz).get();
        }
        stopWatch.stop();
        System.out.println("Container started in " + stopWatch.getTime() + "ms");
        started = true;
    }
}
