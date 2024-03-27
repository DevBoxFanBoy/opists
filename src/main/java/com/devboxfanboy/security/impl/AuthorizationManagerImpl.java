package com.devboxfanboy.security.impl;

import org.casbin.jcasbin.main.Enforcer;

import com.devboxfanboy.exception.AuthorizationException;
import com.devboxfanboy.security.AuthorizationManager;

public class AuthorizationManagerImpl implements AuthorizationManager {
    private Enforcer enforcer;

    public AuthorizationManagerImpl(Enforcer enforcer) {
        this.enforcer = enforcer;
    }

    @Override
    public boolean isAuthorization(String userId, String resourceId, String permission) {
        return enforcer.enforce(userId, resourceId, permission);
    }

    @Override
    public void assertAuthorization(String userId, String resourceId, String permission) throws AuthorizationException {
        if (!isAuthorization(userId, resourceId, permission)) {
            throw new AuthorizationException();
        }
    }
}
