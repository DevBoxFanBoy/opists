import static org.junit.jupiter.api.Assertions.*;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;

import com.devboxfanboy.security.impl.AuthorizationManagerImpl;
import com.devboxfanboy.testframework.Components;
import com.devboxfanboy.testframework.annotation.Container;
import com.devboxfanboy.testframework.extension.ContainerIntegrationTestExtension;
import jakarta.inject.Inject;

@ExtendWith(ContainerIntegrationTestExtension.class)
class AuthorizationManagerImplIntTest {

    @Container
    public static Components components = Components.DEFAULT;

    //@BeforeAll
    //static void setUp() {
    //    FlywayUtil.setUp();
    //}

    @Inject
    private AuthorizationManagerImpl authorizationManagerImpl;

    @Test
    void isAuthorizationWithUnknownUserResourcePermission_notAuthorized() {
        assertFalse(authorizationManagerImpl.isAuthorization("user", "resource", "permission"));
    }

    @Test
    void isAuthorization_isAuthorized() {
        assertTrue(authorizationManagerImpl.isAuthorization("alice", "test", "testpermission"));
    }

}
