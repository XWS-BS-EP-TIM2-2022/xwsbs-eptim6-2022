package com.xws.agentska.model;

import com.fasterxml.jackson.annotation.JsonIgnore;
import org.springframework.security.core.GrantedAuthority;

import javax.persistence.*;

@Entity
@Table(name = "role")
public class Role implements GrantedAuthority {
    private static final long serialVersionUID = 1L;
    public static final String ADMIN_ROLE ="ADMIN_ROLE";
    public static final String USER_ROLE="USER_ROLE";
    public static final String COMPANY_OWNER_ROLE ="COMPANY_OWNER_ROLE";
    @Id
    @Column(name = "id")
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    Long id;
    private String name;
    public Role() {
    }

    public Role(String name) {
        this.name = name;
    }

    @JsonIgnore
    @Override
    public String getAuthority() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getName() {
        return name;
    }

    @JsonIgnore
    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

}
