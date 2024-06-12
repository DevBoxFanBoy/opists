package com.devboxfanboy.persistence.impl;

import org.hibernate.Session;
import org.hibernate.SessionFactory;

import com.devboxfanboy.flyway.FlywayUtil;
import com.devboxfanboy.hibernate.HibernateUtil;
import com.devboxfanboy.persistence.PersistenceIntegrationDataInitializer;

public class PersistenceIntegrationDataInitializerImpl implements PersistenceIntegrationDataInitializer {

    @Override
    public void initialize() {
        FlywayUtil.setUp();
        SessionFactory sessionFactory = HibernateUtil.getSessionFactory();
        Session session = sessionFactory.openSession();
        session.beginTransaction();
        session.createNativeMutationQuery(
                        "insert into opistsschema.CASBIN_RULE (PTYPE,V0,V1,V2,V3,V4,V5,version,ID) values ('p','alice','test','testpermission',NULL,NULL,NULL,0,default)")
                .executeUpdate();
        session.getTransaction().commit();
        session.close();
    }
}
