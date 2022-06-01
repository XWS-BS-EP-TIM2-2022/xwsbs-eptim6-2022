package com.xws.agentska.dto;

import com.xws.agentska.model.JobPosition;
import com.xws.agentska.model.WorkSchedule;
import com.xws.agentska.model.enumerations.ExperienceLevel;

import java.sql.Timestamp;

public class JobOfferDto {
    private JobPosition position;
    private Timestamp validTo;
    private Timestamp createdAt;
    private String description;
    private ExperienceLevel experience;
    private WorkSchedule workSchedule;

    public JobOfferDto() {
    }

    public JobPosition getPosition() {
        return position;
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

    public WorkSchedule getWorkSchedule() {
        return workSchedule;
    }

    public void setWorkSchedule(WorkSchedule workSchedule) {
        this.workSchedule = workSchedule;
    }
}
