package com.devboxfanboy.testframework.testdata;

import static com.devboxfanboy.testframework.testdata.MyServiceIntTest.*;
import static org.mockito.Mockito.*;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.Spy;

import com.devboxfanboy.testframework.Components;
import com.devboxfanboy.testframework.annotation.Container;
import com.devboxfanboy.testframework.extension.ContainerIntegrationTestExtension;
import jakarta.inject.Inject;

@ExtendWith(ContainerIntegrationTestExtension.class)
class MyServiceWithSpyIntTest {

    @Container
    private final Components components = TEST_CONTEXT();

    @Spy
    private SomeDependency someDependency;

    @Inject
    private MyService myService;

    @Test
    void testDoSomething() {
        myService.doSomething();
        verify(someDependency).doSomethingElse();
    }
}
