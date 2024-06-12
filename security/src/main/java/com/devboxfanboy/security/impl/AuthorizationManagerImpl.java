package com.devboxfanboy.security.impl;

import java.io.InputStream;
import java.nio.charset.StandardCharsets;

import org.apache.commons.io.IOUtils;
import org.casbin.jcasbin.main.Enforcer;
import org.casbin.jcasbin.model.Model;
import org.hibernate.SessionFactory;

import com.devboxfanboy.hibernate.HibernateUtil;
import com.devboxfanboy.security.AuthorizationManager;
import com.devboxfanboy.security.exception.AuthorizationException;
import com.devboxfanboy.security.jcasbin.adapter.HibernateAdapter;
import jakarta.annotation.PostConstruct;
import jakarta.ejb.Stateless;
import lombok.AllArgsConstructor;
import lombok.NoArgsConstructor;

@NoArgsConstructor
@AllArgsConstructor
@Stateless
public class AuthorizationManagerImpl implements AuthorizationManager {
    private Enforcer enforcer;

    @PostConstruct
    public void init() {
        SessionFactory sessionFactory = HibernateUtil.getSessionFactory();
        HibernateAdapter adapter = new HibernateAdapter("opistsschema", "casbin_rule", sessionFactory);
        try (InputStream in = AuthorizationManagerImpl.class.getResourceAsStream("/model.conf")) {
            Model model = Model.newModelFromString(IOUtils.toString(in, StandardCharsets.UTF_8));
            enforcer = new Enforcer(model, adapter);
            enforcer.loadPolicy();
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
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
