package com.devboxfanboy;

import org.flywaydb.core.Flyway;
import org.flywaydb.core.api.output.MigrateResult;

public class FlywayUtil {

    public static void setUp() {
        Flyway flyway = Flyway.configure()
                .dataSource(
                        "jdbc:h2:mem:testdb;DB_CLOSE_DELAY=-1;DB_CLOSE_ON_EXIT=FALSE;MODE=PostgreSQL",
                        "dbadmin",
                        "testdb")
                .outOfOrder(true)
                .load();
        MigrateResult migrateResult = flyway.migrate();
        if (!migrateResult.success) {
            throw new RuntimeException("Flyway migration failed");
        }
    }

}
