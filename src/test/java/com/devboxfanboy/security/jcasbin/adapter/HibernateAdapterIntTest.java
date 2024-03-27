package com.devboxfanboy.security.jcasbin.adapter;

import static org.junit.jupiter.api.Assertions.*;

import java.util.Arrays;
import java.util.List;

import org.casbin.jcasbin.main.Enforcer;
import org.hibernate.Session;
import org.hibernate.SessionFactory;
import org.hibernate.Transaction;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import com.devboxfanboy.FlywayUtil;
import com.devboxfanboy.hibernate.HibernateUtil;
import com.devboxfanboy.security.jcasbin.entity.CasbinRule;

class HibernateAdapterIntTest {

    Enforcer enforcer;
    HibernateAdapter adapter;
    SessionFactory sessionFactory;

    @BeforeEach
    void setUp() {
        FlywayUtil.setUp();
        sessionFactory = HibernateUtil.getSessionFactory();
        adapter = new HibernateAdapter("opistsschema", "casbin_rule", sessionFactory);
        enforcer = new Enforcer("src/test/resources/model.conf", adapter);
        enforcer.clearPolicy();
        Transaction tx;
        try (Session session = sessionFactory.openSession()) {
            tx = session.beginTransaction();
            adapter.removeAllPolicies(session);
            tx.commit();
        }
    }

    @AfterEach
    void tearDown() {
        HibernateUtil.shutdown();
    }

    @Test
    void emptySchema() {
        assertEquals("casbin_rule", new HibernateAdapter(null, "casbin_rule", sessionFactory).getSchemaAndTableName());
        assertEquals("casbin_rule", new HibernateAdapter("", "casbin_rule", sessionFactory).getSchemaAndTableName());
        assertEquals("casbin_rule", new HibernateAdapter("casbin_rule", sessionFactory).getSchemaAndTableName());
    }

    @Test
    void notEmptySchema() {
        assertEquals("a.casbin_rule", new HibernateAdapter("a", "casbin_rule", sessionFactory).getSchemaAndTableName());
    }

    @Test
    void withEnforcer() {
        List<List<String>> permissions = enforcer.getPermissionsForUser("alice");
        assertEquals(0, permissions.size());
        enforcer.addPermissionForUser("alice", "data1", "read");
        enforcer.savePolicy();
        permissions = enforcer.getPermissionsForUser("alice");
        assertPermission(permissions, "alice", "data1", "read");
        enforcer.loadPolicy();
        permissions = enforcer.getPermissionsForUser("alice");
        assertPermission(permissions, "alice", "data1", "read");
        enforcer.deletePermissionForUser("alice", "data1", "read");
        permissions = enforcer.getPermissionsForUser("alice");
        assertEquals(0, permissions.size());
    }

    private static void assertPermission(List<List<String>> permissions, String subject, String object, String action) {
        assertEquals(1, permissions.size());
        assertEquals(subject, permissions.get(0).get(0));
        assertEquals(object, permissions.get(0).get(1));
        assertEquals(action, permissions.get(0).get(2));
    }

    @Test
    void loadPolicy() {
        assertEquals(0, getCasbinRules().size());
        assertEquals(0, enforcer.getModel().getPolicy("p", "p").size());
        enforcer.addPolicy("alice", "data1", "read");
        assertEquals(1, enforcer.getModel().getPolicy("p", "p").size());
        enforcer.clearPolicy();
        assertEquals(0, enforcer.getModel().getPolicy("p", "p").size());

        assertDoesNotThrow(() -> adapter.loadPolicy(enforcer.getModel()));

        assertEquals(1, enforcer.getModel().getPolicy("p", "p").size());
        assertEquals(Arrays.asList("alice", "data1", "read"), enforcer.getModel().getPolicy("p", "p").get(0));
        List<CasbinRule> casbinRules = getCasbinRules();
        assertEquals(1, casbinRules.size());
        assertEquals("p", casbinRules.get(0).getPtype());
        assertEquals("alice", casbinRules.get(0).getV0());
        assertEquals("data1", casbinRules.get(0).getV1());
        assertEquals("read", casbinRules.get(0).getV2());
        assertNull(casbinRules.get(0).getV3());
        assertNull(casbinRules.get(0).getV4());
        assertNull(casbinRules.get(0).getV5());
    }

    @Test
    void savePolicy() {
        assertEquals(0, getCasbinRules().size());
        assertEquals(0, enforcer.getModel().getPolicy("p", "p").size());
        enforcer.addPolicy("alice", "data1", "read");
        assertEquals(1, enforcer.getModel().getPolicy("p", "p").size());

        assertDoesNotThrow(() -> adapter.savePolicy(enforcer.getModel()));

        List<CasbinRule> casbinRules = getCasbinRules();
        assertEquals(1, casbinRules.size());
        assertEquals("p", casbinRules.get(0).getPtype());
        assertEquals("alice", casbinRules.get(0).getV0());
        assertEquals("data1", casbinRules.get(0).getV1());
        assertEquals("read", casbinRules.get(0).getV2());
        assertNull(casbinRules.get(0).getV3());
        assertNull(casbinRules.get(0).getV4());
        assertNull(casbinRules.get(0).getV5());
    }

    @Test
    void savePolicyClearedPolicy_noPolicyStored() {
        assertEquals(0, getCasbinRules().size());
        assertEquals(0, enforcer.getModel().getPolicy("p", "p").size());
        enforcer.addPolicy("alice", "data1", "read");
        assertEquals(1, enforcer.getModel().getPolicy("p", "p").size());
        enforcer.clearPolicy();

        assertDoesNotThrow(() -> adapter.savePolicy(enforcer.getModel()));

        assertEquals(0, getCasbinRules().size());
    }

    @Test
    void addPolicy() {
        assertEquals(0, getCasbinRules().size());
        assertEquals(0, enforcer.getModel().getPolicy("p", "p").size());

        assertDoesNotThrow(() -> adapter.addPolicy("p", "p", List.of("alice", "data1", "read")));

        List<CasbinRule> casbinRules = getCasbinRules();
        assertEquals(1, casbinRules.size());
        assertEquals("p", casbinRules.get(0).getPtype());
        assertEquals("alice", casbinRules.get(0).getV0());
        assertEquals("data1", casbinRules.get(0).getV1());
        assertEquals("read", casbinRules.get(0).getV2());
        assertNull(casbinRules.get(0).getV3());
        assertNull(casbinRules.get(0).getV4());
        assertNull(casbinRules.get(0).getV5());
    }

    @Test
    void addPolicyAlreadyAdded_notStoredTwice() {
        assertEquals(0, getCasbinRules().size());
        assertEquals(0, enforcer.getModel().getPolicy("p", "p").size());

        assertDoesNotThrow(() -> adapter.addPolicy("p", "p", List.of("alice", "data1", "read")));
        assertDoesNotThrow(() -> adapter.addPolicy("p", "p", List.of("alice", "data1", "read")));

        List<CasbinRule> casbinRules = getCasbinRules();
        assertEquals(1, casbinRules.size());
        assertEquals("p", casbinRules.get(0).getPtype());
        assertEquals("alice", casbinRules.get(0).getV0());
        assertEquals("data1", casbinRules.get(0).getV1());
        assertEquals("read", casbinRules.get(0).getV2());
        assertNull(casbinRules.get(0).getV3());
        assertNull(casbinRules.get(0).getV4());
        assertNull(casbinRules.get(0).getV5());
    }

    @Test
    void removePolicy() {
        assertEquals(0, getCasbinRules().size());
        assertEquals(0, enforcer.getModel().getPolicy("p", "p").size());
        enforcer.addPolicy("alice", "data1", "read");
        assertEquals(1, enforcer.getModel().getPolicy("p", "p").size());

        assertDoesNotThrow(() -> adapter.removePolicy("p", "p", List.of("alice", "data1", "read")));

        assertEquals(0, getCasbinRules().size());

        enforcer.clearPolicy();
        enforcer.loadPolicy();
        assertEquals(0, enforcer.getModel().getPolicy("p", "p").size());
    }

    //CasbinRule(id=null, version=0, ptype=p, v0=alice, v1=data1, v2=read, v3=null, v4=null, v5=null)
    @Test
    void removeFilteredPolicy() {
        assertEquals(0, getCasbinRules().size());
        assertEquals(0, enforcer.getModel().getPolicy("p", "p").size());
        enforcer.addPolicy("alice", "data1", "read");
        assertEquals(1, enforcer.getModel().getPolicy("p", "p").size());
        assertEquals(1, getCasbinRules().size());

        assertDoesNotThrow(() -> adapter.removeFilteredPolicy("p", "p", 0, "alice"));

        assertEquals(0, getCasbinRules().size());
    }

    private List<CasbinRule> getCasbinRules() {
        List<CasbinRule> casbinRules;
        Transaction tx;
        try (Session session = sessionFactory.openSession()) {
            tx = session.beginTransaction();
            casbinRules = session.createQuery("from CasbinRule", CasbinRule.class).list();
            tx.commit();
        }
        return casbinRules;
    }
}
