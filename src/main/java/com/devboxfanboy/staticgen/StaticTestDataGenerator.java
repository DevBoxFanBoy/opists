package com.devboxfanboy.staticgen;

import org.apache.commons.lang3.ArrayUtils;
import org.casbin.jcasbin.main.Enforcer;
import org.hibernate.SessionFactory;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import ch.qos.logback.classic.Level;
import ch.qos.logback.classic.LoggerContext;
import ch.qos.logback.core.util.StatusPrinter;
import com.devboxfanboy.FlywayUtil;
import com.devboxfanboy.hibernate.HibernateUtil;
import com.devboxfanboy.security.jcasbin.adapter.HibernateAdapter;

public class StaticTestDataGenerator {

    private static final Logger LOGGER = LoggerFactory.getLogger(StaticTestDataGenerator.class);

    static Enforcer enforcer;
    static HibernateAdapter adapter;
    static SessionFactory sessionFactory;

    public static void main(String[] args) {
        // assume SLF4J is bound to logback in the current environment
        LoggerContext lc = (LoggerContext) LoggerFactory.getILoggerFactory();
        // print logback's internal status
        StatusPrinter.print(lc);

        if (ArrayUtils.contains(args, "disableSQLLogging")) {
            ((ch.qos.logback.classic.Logger) LoggerFactory.getLogger("jdbc.sqlonly")).setLevel(Level.ERROR);
        }

        if (LOGGER.isInfoEnabled()) {
            LOGGER.info("jdbc.sqlonly is enabled for %s".formatted(lc.getLogger("jdbc.sqlonly").getLevel()));
        }

        LOGGER.info("Setting up test data...");
        FlywayUtil.setUp();
        sessionFactory = HibernateUtil.getSessionFactory("hibernate-gentestdata.cfg.xml");
        adapter = new HibernateAdapter("opistsschema", "casbin_rule", sessionFactory);
        enforcer = new Enforcer("src/test/resources/model.conf", adapter);
        LOGGER.info("Inserting test data...");
        enforcer.addPermissionForUser("alice", "data1", "read");
        enforcer.addPermissionForUser("alice", "data2", "write");
        LOGGER.info("Done test data...");

    }
}
