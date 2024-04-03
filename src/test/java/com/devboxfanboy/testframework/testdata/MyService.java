package com.devboxfanboy.testframework.testdata;

import jakarta.annotation.PostConstruct;
import jakarta.enterprise.context.ApplicationScoped;
import jakarta.inject.Inject;

@ApplicationScoped
public class MyService {

    @Inject
    private SomeDependency someDependency;

    @PostConstruct
    public void init() {
        System.out.println("init MyService");
    }

    public void doSomething() {
        someDependency.doSomethingElse();
    }
}
