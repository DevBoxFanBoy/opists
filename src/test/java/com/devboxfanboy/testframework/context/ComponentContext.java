package com.devboxfanboy.testframework.context;

import java.util.HashMap;
import java.util.Map;

import com.devboxfanboy.config.feature.Feature;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.ToString;


@NoArgsConstructor
@Getter
@ToString
@EqualsAndHashCode
public class ComponentContext {
    private final Map<Feature, Boolean> featureFlags = new HashMap<>();

    public boolean isFeatureEnabled(Feature feature) {
        return featureFlags.getOrDefault(feature, false);
    }

    public ComponentContext addFeatureFlag(Feature feature) {
        featureFlags.put(feature, true);
        return this;
    }

}
