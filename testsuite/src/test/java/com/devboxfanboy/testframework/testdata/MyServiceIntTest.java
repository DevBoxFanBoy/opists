package com.devboxfanboy.testframework.testdata;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;

import com.devboxfanboy.feature.Feature;
import com.devboxfanboy.testframework.Components;
import com.devboxfanboy.testframework.annotation.Container;
import com.devboxfanboy.testframework.config.ComponentConfiguration;
import com.devboxfanboy.testframework.context.ComponentContext;
import com.devboxfanboy.testframework.extension.ContainerIntegrationTestExtension;
import jakarta.inject.Inject;

@ExtendWith(ContainerIntegrationTestExtension.class)
class MyServiceIntTest {

    @Container
    private final Components components = TEST_CONTEXT();

    public static Components TEST_CONTEXT() {
        return Components.newContext()
                .configuration(ComponentConfiguration.builder().addClass(
                                com.devboxfanboy.testframework.testdata.MyService.class).addClass(SomeDependency.class)
                        .build().init(new ComponentContext().addFeatureFlag(Feature.NONE))).build();
    }

    @Inject
    private MyService myService;

    @Test
    void testDoSomething() {
        myService.doSomething();
    }

}
