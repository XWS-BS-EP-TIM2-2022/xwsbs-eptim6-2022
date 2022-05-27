package com.xws.agentska.model;

import com.xws.agentska.model.enumerations.Status;

import javax.persistence.*;

@Entity
public class JobPosition {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    public String name;
    public boolean confirmed;

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

}
