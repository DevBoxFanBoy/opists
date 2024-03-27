package com.devboxfanboy.security.jcasbin.adapter;

import java.util.List;
import java.util.Map;

import org.apache.commons.collections4.CollectionUtils;
import org.apache.commons.lang3.StringUtils;
import org.casbin.jcasbin.model.Assertion;
import org.casbin.jcasbin.model.Model;
import org.casbin.jcasbin.persist.Adapter;
import org.casbin.jcasbin.persist.Helper;
import org.hibernate.Session;
import org.hibernate.SessionFactory;
import org.hibernate.Transaction;
import org.hibernate.query.MutationQuery;
import org.hibernate.query.Query;

import com.devboxfanboy.security.jcasbin.entity.CasbinRule;
import lombok.Getter;

@Getter
public class HibernateAdapter implements Adapter {
    private static final String PTYPE = "ptype";
    private static final String DELETE_FROM_STR = "delete from %s";
    private final SessionFactory sessionFactory;

    private final String schemaAndTableName;

    public HibernateAdapter(String tableName, SessionFactory sessionFactory) {
        this(null, tableName, sessionFactory);
    }

    public HibernateAdapter(String schemaName, String tableName, SessionFactory sessionFactory) {
        this.sessionFactory = sessionFactory;
        if (StringUtils.isNotEmpty(schemaName)) {
            this.schemaAndTableName = StringUtils.joinWith(".", schemaName, tableName);
        } else {
            this.schemaAndTableName = tableName;
        }
    }

    @SuppressWarnings("JpaQlInspection")
    @Override
    public void loadPolicy(Model model) {
        Transaction tx;
        List<CasbinRule> casbinRules;
        try (Session session = sessionFactory.openSession()) {
            tx = session.beginTransaction();
            casbinRules = session.createQuery("from CasbinRule",
                    CasbinRule.class).list();
            for (CasbinRule casbinRule : casbinRules) {
                loadOnePolicyLine(casbinRule, model);
            }
            tx.commit();
        }
    }

    private static void loadOnePolicyLine(CasbinRule line, Model model) {
        String lineText = line.getPtype();
        if (line.getV0() != null) {
            lineText = lineText + ", " + line.getV0();
        }

        if (line.getV1() != null) {
            lineText = lineText + ", " + line.getV1();
        }

        if (line.getV2() != null) {
            lineText = lineText + ", " + line.getV2();
        }

        if (line.getV3() != null) {
            lineText = lineText + ", " + line.getV3();
        }

        if (line.getV4() != null) {
            lineText = lineText + ", " + line.getV4();
        }

        if (line.getV5() != null) {
            lineText = lineText + ", " + line.getV5();
        }

        Helper.loadPolicyLine(lineText, model);
    }

    @SuppressWarnings({ "JpaQlInspection", "SqlNoDataSourceInspection" })
    @Override
    public void savePolicy(Model model) {
        Transaction tx;
        try (Session session = this.sessionFactory.openSession()) {
            tx = session.beginTransaction();
            List<CasbinRule> casbinRules = session.createQuery("from CasbinRule",
                    CasbinRule.class).list();
            if (!CollectionUtils.isEmpty(casbinRules)) {
                removeAllPolicies(session);
            }
            // Iterate over the policy in the model and persist each rule
            for (Map.Entry<String, Assertion> entry : model.model.get("p").entrySet()) {
                String ptype = entry.getKey();
                for (List<String> rule : entry.getValue().policy) {
                    CasbinRule line = this.toCasbinRule(ptype, rule);
                    session.persist(line);
                }
            }
            tx.commit();
        }
    }

    protected void removeAllPolicies(Session session) {
        session.createNativeMutationQuery(DELETE_FROM_STR.formatted(schemaAndTableName)).executeUpdate();
    }

    @Override
    public void addPolicy(String sec, String ptype, List<String> rule) {
        if (!CollectionUtils.isEmpty(rule)) {
            CasbinRule line = this.toCasbinRule(ptype, rule);
            try (Session session = this.sessionFactory.openSession()) {
                Transaction tx = session.beginTransaction();
                if (!ruleExists(session, line)) {
                    session.persist(line);
                    tx.commit();
                }
            }
        }
    }

    private CasbinRule toCasbinRule(String ptype, List<String> rule) {
        CasbinRule line = new CasbinRule();
        line.setPtype(ptype);
        if (!rule.isEmpty()) {
            line.setV0(rule.get(0));
        }

        if (rule.size() > 1) {
            line.setV1(rule.get(1));
        }

        if (rule.size() > 2) {
            line.setV2(rule.get(2));
        }

        if (rule.size() > 3) {
            line.setV3(rule.get(3));
        }

        if (rule.size() > 4) {
            line.setV4(rule.get(4));
        }

        if (rule.size() > 5) {
            line.setV5(rule.get(5));
        }

        return line;
    }

    @SuppressWarnings({ "JpaQlInspection", "SqlNoDataSourceInspection" })
    protected boolean ruleExists(Session session, CasbinRule rule) {
        String selectHqlString = getSelectHqlString(rule);

        Query<CasbinRule> query = session.createQuery(selectHqlString, CasbinRule.class);
        setParameterOrIsNullCheck(rule, query);

        List<CasbinRule> existingRules = query.list();
        return !existingRules.isEmpty();
    }

    @SuppressWarnings("DuplicatedCode")
    private static void setParameterOrIsNullCheck(CasbinRule rule, Query<CasbinRule> query) {
        query.setParameter(PTYPE, rule.getPtype());
        if (rule.getV0() != null) {
            query.setParameter("v0", rule.getV0());
        }
        if (rule.getV1() != null) {
            query.setParameter("v1", rule.getV1());
        }
        if (rule.getV2() != null) {
            query.setParameter("v2", rule.getV2());
        }
        if (rule.getV3() != null) {
            query.setParameter("v3", rule.getV3());
        }
        if (rule.getV4() != null) {
            query.setParameter("v4", rule.getV4());
        }
        if (rule.getV5() != null) {
            query.setParameter("v5", rule.getV5());
        }
    }

    private static String getSelectHqlString(CasbinRule rule) {
        StringBuilder queryBuilder = new StringBuilder("from CasbinRule where ptype = :ptype");
        return getParameterOrIsNullStrings(rule, queryBuilder);
    }

    private static String getParameterOrIsNullStrings(CasbinRule rule, StringBuilder queryBuilder) {
        if (rule.getV0() != null) {
            queryBuilder.append(" and v0 = :v0");
        } else {
            queryBuilder.append(" and v0 is null");
        }
        if (rule.getV1() != null) {
            queryBuilder.append(" and v1 = :v1");
        } else {
            queryBuilder.append(" and v1 is null");
        }
        if (rule.getV2() != null) {
            queryBuilder.append(" and v2 = :v2");
        } else {
            queryBuilder.append(" and v2 is null");
        }
        if (rule.getV3() != null) {
            queryBuilder.append(" and v3 = :v3");
        } else {
            queryBuilder.append(" and v3 is null");
        }
        if (rule.getV4() != null) {
            queryBuilder.append(" and v4 = :v4");
        } else {
            queryBuilder.append(" and v4 is null");
        }
        if (rule.getV5() != null) {
            queryBuilder.append(" and v5 = :v5");
        } else {
            queryBuilder.append(" and v5 is null");
        }
        return queryBuilder.toString();
    }

    @SuppressWarnings({ "JpaQlInspection", "SqlNoDataSourceInspection" })
    @Override
    public void removePolicy(String sec, String ptype, List<String> rule) {
        if (!CollectionUtils.isEmpty(rule)) {
            CasbinRule line = this.toCasbinRule(ptype, rule);
            Transaction tx;
            try (Session session = this.sessionFactory.openSession()) {
                tx = session.beginTransaction();
                String deleteHqlString = getDeleteHqlString(line);
                MutationQuery deleteQuery = session.createNativeMutationQuery(deleteHqlString);
                //noinspection DuplicatedCode
                deleteQuery.setParameter(PTYPE, line.getPtype());
                if (line.getV0() != null) {
                    deleteQuery.setParameter("v0", line.getV0());
                }
                if (line.getV1() != null) {
                    deleteQuery.setParameter("v1", line.getV1());
                }
                if (line.getV2() != null) {
                    deleteQuery.setParameter("v2", line.getV2());
                }
                if (line.getV3() != null) {
                    deleteQuery.setParameter("v3", line.getV3());
                }
                if (line.getV4() != null) {
                    deleteQuery.setParameter("v4", line.getV4());
                }
                if (line.getV5() != null) {
                    deleteQuery.setParameter("v5", line.getV5());
                }
                int deletedRows = deleteQuery.executeUpdate();
                if (deletedRows > 0) {
                    tx.commit();
                }
            }
        }
    }

    private String getDeleteHqlString(CasbinRule rule) {
        StringBuilder queryBuilder = new StringBuilder(
                DELETE_FROM_STR.formatted(schemaAndTableName) + " where ptype = :ptype");
        return getParameterOrIsNullStrings(rule, queryBuilder);
    }

    @SuppressWarnings({ "JpaQlInspection", "SqlNoDataSourceInspection" })
    @Override
    public void removeFilteredPolicy(String sec, String ptype, int fieldIndex, String... fieldValues) {
        if (fieldValues.length > 0) {
            StringBuilder queryString = new StringBuilder(
                    DELETE_FROM_STR.formatted(schemaAndTableName) + " where ptype = :ptype");
            String[] fields = { "v0", "v1", "v2", "v3", "v4", "v5" };
            for (int i = 0; i < fieldValues.length; i++) {
                queryString.append(" and ").append(fields[fieldIndex + i]).append(" = :")
                        .append(fields[fieldIndex + i]);
            }
            Transaction tx;
            MutationQuery deleteQuery;
            try (Session session = this.sessionFactory.openSession()) {
                tx = session.beginTransaction();
                deleteQuery = session.createNativeMutationQuery(
                        queryString.toString());
                deleteQuery.setParameter(PTYPE, ptype);
                for (int i = 0; i < fieldValues.length; i++) {
                    deleteQuery.setParameter(fields[fieldIndex + i], fieldValues[i]);
                }

                int deletedRows = deleteQuery.executeUpdate();
                if (deletedRows > 0) {
                    tx.commit();
                }
            }
        }
    }
}
