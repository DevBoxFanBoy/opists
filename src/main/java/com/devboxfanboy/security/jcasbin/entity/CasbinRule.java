package com.devboxfanboy.security.jcasbin.entity;

import jakarta.persistence.Entity;
import jakarta.persistence.Id;
import jakarta.persistence.Table;
import lombok.AccessLevel;
import lombok.AllArgsConstructor;
import lombok.EqualsAndHashCode;
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
@Getter
@Setter
@EqualsAndHashCode
@ToString
@AllArgsConstructor(access = AccessLevel.PACKAGE)
@NoArgsConstructor()
@Entity
@Table(name = "casbin_rule", schema = "opistsschema")
public class CasbinRule {

    @Id
    private Long id;
    private int version;

    private String ptype;
    private String v0;
    private String v1;
    private String v2;
    private String v3;
    private String v4;
    private String v5;
}
