package com.devboxfanboy.testframework.container.impl;

import com.devboxfanboy.testframework.container.TestContainer;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.experimental.Delegate;

@SuppressWarnings("ClassCanBeRecord")
@Data
@AllArgsConstructor
public class IntegrationTestContainer implements TestContainer {

    @Delegate(types = TestContainer.class)
    private final TestContainer testContainer;
}
