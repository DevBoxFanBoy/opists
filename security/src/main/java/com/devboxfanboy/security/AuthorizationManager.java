package com.devboxfanboy.security;

import com.devboxfanboy.security.exception.AuthorizationException;

public interface AuthorizationManager {
    boolean isAuthorization(String userId, String resourceId, String permission);
    void assertAuthorization(String userId, String resourceId, String permission)
            throws AuthorizationException;
}
