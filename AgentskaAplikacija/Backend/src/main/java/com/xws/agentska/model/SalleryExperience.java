package com.xws.agentska.model;

import javax.persistence.Entity;

@Entity
public class SalleryExperience extends UserExperience {
    private double sallery;

    public SalleryExperience() {
    }

    public double getSallery() {
        return sallery;
    }

    public void setSallery(double sallery) {
        this.sallery = sallery;
    }
}
