package com.devboxfanboy.security.impl;

import static org.junit.jupiter.api.Assertions.*;

import org.casbin.jcasbin.main.Enforcer;
import org.casbin.jcasbin.persist.Adapter;
import org.hibernate.SessionFactory;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import com.devboxfanboy.FlywayUtil;
import com.devboxfanboy.exception.AuthorizationException;
import com.devboxfanboy.hibernate.HibernateUtil;
import com.devboxfanboy.security.jcasbin.adapter.HibernateAdapter;

class AuthorizationManagerIntTest {

    Enforcer enforcer;

    AuthorizationManagerImpl authorizationManager;

    @BeforeEach
    void setUp() {
        enforcer = new Enforcer("src/test/resources/model.conf", "src/test/resources/policy.csv");
        enforcer.addPermissionForUser("alice", "data1", "write");
        authorizationManager = new AuthorizationManagerImpl(enforcer);
    }

    @Test
    void authorizeAliceReadData1() {
        assertTrue(authorizationManager.isAuthorization("alice", "data1", "read"));
    }

    @Test
    void authorizeAliceWriteData1() {
        assertTrue(authorizationManager.isAuthorization("alice", "data1", "write"));
    }

    @Test
    void authorizeAliceReadData2() {
        assertFalse(authorizationManager.isAuthorization("alice", "data2", "read"));
    }

    @Test
    void assertAuthorizationAliceWriteData1NotThrowAuthorizationException() {
        assertDoesNotThrow(() -> authorizationManager.assertAuthorization("alice", "data1", "write"));
    }

    @Test
    void assertAuthorizationAliceReadData1ThrowAuthorizationException() {
        assertThrows(AuthorizationException.class,
                () -> authorizationManager.assertAuthorization("alice", "throooww", "read"));
    }

    @Test
    void testHibernateAdapter() {
        FlywayUtil.setUp();
        SessionFactory sessionFactory = HibernateUtil.getSessionFactory();
        Adapter adapter = new HibernateAdapter("opistsschema", "casbin_rule", sessionFactory);
        enforcer = new Enforcer("src/test/resources/model.conf", adapter);
        enforcer.clearPolicy();
        authorizationManager = new AuthorizationManagerImpl(enforcer);
        assertFalse(authorizationManager.isAuthorization("alice", "data1", "read"));
        enforcer.addPermissionForUser("alice", "data1", "read");
        assertTrue(authorizationManager.isAuthorization("alice", "data1", "read"));
    }

}
