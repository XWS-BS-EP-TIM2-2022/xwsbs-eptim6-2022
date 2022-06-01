package com.xws.agentska.model;

import javax.persistence.Embeddable;

@Embeddable
public class ContactInfo {
    private String email;
    private String phone;
    public ContactInfo(){}

    public ContactInfo(String email, String phone) {
        this.email = email;
        this.phone = phone;
    }

    public String getPhone() {
        return phone;
    }

    public void setPhone(String phone) {
        this.phone = phone;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }
}