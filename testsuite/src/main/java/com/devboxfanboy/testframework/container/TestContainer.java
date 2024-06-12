package com.devboxfanboy.testframework.container;

public interface TestContainer {

    Object get(Class<?> clazz);

    void shutdown();
    void start();
}
