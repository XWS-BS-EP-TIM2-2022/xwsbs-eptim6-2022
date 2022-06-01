package com.xws.agentska.dto;

public class UserTokenState {
    private String jwt;
    private String role;
    private Long id;

    public UserTokenState(String jwt, String role, Long id) {
        this.jwt = jwt;
        this.role = role;
        this.id = id;
    }

    public String getJwt() {
        return jwt;
    }

    public void setJwt(String jwt) {
        this.jwt = jwt;
    }

    public String getRole() {
        return role;
    }

    public void setRole(String role) {
        this.role = role;
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }
}
