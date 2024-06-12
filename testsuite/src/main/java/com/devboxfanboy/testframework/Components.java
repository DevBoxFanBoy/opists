package com.devboxfanboy.testframework;

import com.devboxfanboy.feature.Feature;
import com.devboxfanboy.testframework.config.ComponentConfiguration;
import com.devboxfanboy.testframework.context.ComponentContext;
import lombok.Builder;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.ToString;

@Builder(builderMethodName = "newContext")
@Getter
@ToString
@EqualsAndHashCode
public class Components {
    private ComponentConfiguration configuration;

    public static Components DEFAULT = of(new ComponentContext().addFeatureFlag(Feature.ALL));

    protected static Components of(ComponentContext context) {
        return Components.newContext()
                .configuration(ComponentConfiguration.builder().build().init(context))
                .build();
    }
}
