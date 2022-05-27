package com.xws.agentska.model;


import com.xws.agentska.model.enumerations.ExperienceLevel;

import javax.persistence.*;
import java.sql.Timestamp;

@Entity
public class JobOffer {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    @ManyToOne
    @JoinColumn(name = "position_id")
    private JobPosition position;
    private Timestamp validTo;
    private Timestamp createdAt;
    private String description;
    private ExperienceLevel experience;
    @Embedded
    private WorkSchedule workSchedule;

    public JobOffer() {
    }

    public void setPosition(JobPosition position) {
        this.position = position;
    }

    public Timestamp getValidTo() {
        return validTo;
    }

    public void setValidTo(Timestamp validTo) {
        this.validTo = validTo;
    }

    public Timestamp getCreatedAt() {
        return createdAt;
    }

    public void setCreatedAt(Timestamp createdAt) {
        this.createdAt = createdAt;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public ExperienceLevel getExperience() {
        return experience;
    }

    public void setExperience(ExperienceLevel experience) {
        this.experience = experience;
    }

    public void setWorkSchedule(WorkSchedule workSchedule) {
        this.workSchedule = workSchedule;
    }

    public JobPosition getPosition() {
        return position;
    }

    public WorkSchedule getWorkSchedule() {
        return workSchedule;
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }
}
