package com.devboxfanboy.hibernate;

import org.apache.commons.lang3.StringUtils;
import org.hibernate.SessionFactory;
import org.hibernate.boot.Metadata;
import org.hibernate.boot.MetadataSources;
import org.hibernate.boot.registry.StandardServiceRegistry;
import org.hibernate.boot.registry.StandardServiceRegistryBuilder;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import lombok.AccessLevel;
import lombok.NoArgsConstructor;

@NoArgsConstructor(access = AccessLevel.PRIVATE)
public class HibernateUtil {

    private static final Logger LOGGER = LoggerFactory.getLogger(HibernateUtil.class);

    private static StandardServiceRegistry registry;
    private static SessionFactory sessionFactory;

    public static SessionFactory getSessionFactory() {
        return getSessionFactory(null);
    }

    public static SessionFactory getSessionFactory(String hibernateConfigFile) {
        if (sessionFactory == null) {
            try {
                // Create registry
                if (StringUtils.isNotEmpty(hibernateConfigFile)) {
                    registry = new StandardServiceRegistryBuilder().configure(hibernateConfigFile).build();
                } else {
                    registry = new StandardServiceRegistryBuilder().configure().build();
                }
                // Create MetadataSources
                MetadataSources sources = new MetadataSources(registry);

                // Create Metadata
                Metadata metadata = sources.getMetadataBuilder().build();

                // Create SessionFactory
                sessionFactory = metadata.getSessionFactoryBuilder().build();

            } catch (Exception e) {
                LOGGER.error(e.getLocalizedMessage(), e);
                if (registry != null) {
                    StandardServiceRegistryBuilder.destroy(registry);
                }
            }
        }
        return sessionFactory;
    }

    public static void shutdown() {
        if (registry != null) {
            StandardServiceRegistryBuilder.destroy(registry);
        }
    }

}
