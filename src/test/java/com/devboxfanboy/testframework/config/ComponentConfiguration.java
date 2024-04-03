package com.devboxfanboy.testframework.config;

import java.util.HashSet;
import java.util.Set;

import com.devboxfanboy.config.feature.Feature;
import com.devboxfanboy.security.impl.AuthorizationManagerImpl;
import com.devboxfanboy.testframework.context.ComponentContext;
import lombok.AccessLevel;
import lombok.AllArgsConstructor;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.ToString;

@NoArgsConstructor
@AllArgsConstructor(access = AccessLevel.PRIVATE)
@Getter
@ToString
@EqualsAndHashCode
public class ComponentConfiguration {

    private Set<Class<?>> classes = new HashSet<>();

    public static ComponentConfigurationBuilder builder() {
        return new ComponentConfigurationBuilder();
    }

    public ComponentConfiguration init(ComponentContext context) {
        if (context.isFeatureEnabled(Feature.ALL)) {
            classes.add(AuthorizationManagerImpl.class);
        }
        return this;
    }

    public static class ComponentConfigurationBuilder {

        private Set<Class<?>> classes = new HashSet<>();

        public ComponentConfigurationBuilder addClass(Class<?> clazz) {
            classes.add(clazz);
            return this;
        }

        public ComponentConfiguration build() {
            return new ComponentConfiguration(classes);
        }
    }
}
