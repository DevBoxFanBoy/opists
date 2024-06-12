package com.devboxfanboy.security.jcasbin.entity;

import java.util.Objects;

import org.hibernate.annotations.JdbcTypeCode;
import org.hibernate.proxy.HibernateProxy;
import org.hibernate.type.SqlTypes;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.Table;
import jakarta.persistence.Version;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;
import lombok.ToString;

/**
 * Represents a rule in the Casbin framework.
 * <p>
 * The fields v0 to v5 represent elements of an access control rule and their meaning can vary depending on the context.
 * Here are some examples for different models:
 * <ul>
 *     <li>ACL (Access Control List) Model:
 *         <ul>
 *             <li>v0: Subject (User)</li>
 *             <li>v1: Object (Resource)</li>
 *             <li>v2: Action (read, write, delete, etc.)</li>
 *         </ul>
 *     </li>
 *     <li>RBAC (Role-Based Access Control) Model:
 *         <ul>
 *             <li>v0: Subject (Role)</li>
 *             <li>v1: Object (Resource)</li>
 *             <li>v2: Action (read, write, delete, etc.)</li>
 *         </ul>
 *     </li>
 *     <li>ABAC (Attribute-Based Access Control) Model:
 *         <ul>
 *             <li>v0: Subject with attributes</li>
 *             <li>v1: Object with attributes</li>
 *             <li>v2: Action with attributes</li>
 *             <li>v3 to v5: Additional attributes</li>
 *         </ul>
 *     </li>
 *     <li>RESTful Model:
 *         <ul>
 *             <li>v0: Subject (User)</li>
 *             <li>v1: Path (Resource)</li>
 *             <li>v2: Method (GET, POST, PUT, DELETE, etc.)</li>
 *         </ul>
 *     </li>
 * </ul>
 * <p>
 * Please note that the interpretation of v0 to v5 depends on your specific Casbin model configuration
 * and the requirements of your application. It's possible that v0 to v5 have a different meaning in your application.
 */
@NoArgsConstructor
@Getter
@Setter
@ToString
@Entity
@Table(name = "casbin_rule", schema = "opistsschema")
public class CasbinRule {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "id", nullable = false)
    private Long id;

    @Version
    @Column(name = "version")
    @JdbcTypeCode(SqlTypes.INTEGER)
    private Integer version;

    @Column(name = "ptype")
    private String ptype;
    @Column(name = "v0")
    private String v0;
    @Column(name = "v1")
    private String v1;
    @Column(name = "v2")
    private String v2;
    @Column(name = "v3")
    private String v3;
    @Column(name = "v4")
    private String v4;
    @Column(name = "v5")
    private String v5;

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null) return false;
        Class<?> oEffectiveClass;
        if (o instanceof HibernateProxy hibernateProxy) {
            oEffectiveClass = hibernateProxy.getHibernateLazyInitializer().getPersistentClass();
        } else {
            oEffectiveClass = o.getClass();
        }
        Class<?> thisEffectiveClass;
        if (this instanceof HibernateProxy hibernateProxy) {
            thisEffectiveClass = hibernateProxy.getHibernateLazyInitializer().getPersistentClass();
        } else {
            thisEffectiveClass = this.getClass();
        }
        if (thisEffectiveClass != oEffectiveClass) {
            return false;
        }
        CasbinRule that = (CasbinRule) o;
        return getId() != null && Objects.equals(getId(), that.getId());
    }

    @Override
    public int hashCode() {
        if (this instanceof HibernateProxy hibernateProxy) {
            return hibernateProxy.getHibernateLazyInitializer().getPersistentClass().hashCode();
        }
        return getClass().hashCode();
    }
}
