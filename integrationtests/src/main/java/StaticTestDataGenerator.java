import java.util.List;

import org.apache.commons.lang3.ArrayUtils;
import org.casbin.jcasbin.main.Enforcer;
import org.hibernate.SessionFactory;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import ch.qos.logback.classic.Level;
import ch.qos.logback.classic.LoggerContext;
import ch.qos.logback.core.util.StatusPrinter;
import com.devboxfanboy.flyway.FlywayUtil;
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
        enforcer = new Enforcer("src/main/resources/abac-model.conf", adapter);
        LOGGER.info("Inserting test data...");
        // list of users
        List<String> listOfUsernames = List.of("awesome","alice", "bob", "dude", "guy");

        // add policies structurally
        enforcer.addPolicy("admin", "resources", "manage");
        enforcer.addPolicy("user", "own_resources", "manage");
        enforcer.addPolicy("project_admin", "project_resources", "manage");
        enforcer.addPolicy("project_member", "assigned_issues", "manage");
        enforcer.addPolicy("process_admin", "process_elements", "manage");
        enforcer.addPolicy("view_admin", "filters", "manage");
        enforcer.addPolicy("workflow_admin", "issue_status", "manage");
        enforcer.addPolicy("user", "own_resources", "share");
        enforcer.addPolicy("admin", "membership", "restrict");
        enforcer.addPolicy("system_admin", "administrators", "appoint_remove");
        enforcer.addPolicy("system_admin", "system", "configure");
        enforcer.addPolicy("system_admin", "users", "create_modify_delete");
        enforcer.addPolicy("user", "own_profile", "modify");
        int userId = 0;
        for (String username : listOfUsernames) {
            enforcer.addPolicy(username, "user:%d.profile".formatted(userId), "modify"); //user can modify his profile
            ++userId;
        }
        // add system-admin
        enforcer.addPolicy("awesome", "*", "system-admin"); //awesome is a system-admin

        // add policies for users
        enforcer.addPolicy("alice", "project:1", "manage"); //project_admin of alice-project
        enforcer.addPolicy("bob", "project:1", "member"); //bob is a member of alice-project

        // bob share issue 123 with guy
        enforcer.addPolicy("bob", "issue:123", "read");  //project_admin give bob the right to read
        enforcer.addPolicy("bob", "issue:123", "share"); //project_admin give bob the right to share
        enforcer.addPolicy("guy", "issue:123", "read");  //bob share the right to read with guy


        LOGGER.info("Done test data...");

    }
}
