package com.devboxfanboy.testframework.testdata;

import jakarta.enterprise.context.ApplicationScoped;

@ApplicationScoped
public class SomeDependency {

    public void doSomethingElse() {
        System.out.println("Doing something else");
    }
}
