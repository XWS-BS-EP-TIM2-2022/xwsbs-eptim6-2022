package com.xws.agentska.dto;

import com.xws.agentska.model.JobOffer;

public class JobOfferDislinktDto {
    private String position;
    private String description;
    private String experience;
    private String createdOn;
    private String validUntil;
    private String workScheduleTitle;
    private String workScheduleHours;
    private String jobOfferUrl;

    public JobOfferDislinktDto() {
    }
    public JobOfferDislinktDto(JobOffer offer) {
        this.position=offer.getPosition().getName();
        this.description=offer.getDescription();
        this.experience=offer.getExperience().name();
        this.createdOn=offer.getCreatedAt().toString();
        this.validUntil=offer.getValidTo().toString();
        this.workScheduleTitle=offer.getWorkSchedule().getName();
        this.workScheduleHours= String.valueOf(offer.getWorkSchedule().getHoursPerWeek());
        this.jobOfferUrl="http://localhost";//TODO: Implementirati
    }


    public String getPosition() {
        return position;
    }

    public void setPosition(String position) {
        this.position = position;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public String getExperience() {
        return experience;
    }

    public void setExperience(String experience) {
        this.experience = experience;
    }

    public String getCreatedOn() {
        return createdOn;
    }

    public void setCreatedOn(String createdOn) {
        this.createdOn = createdOn;
    }

    public String getValidUntil() {
        return validUntil;
    }

    public void setValidUntil(String validUntil) {
        this.validUntil = validUntil;
    }

    public String getWorkScheduleTitle() {
        return workScheduleTitle;
    }

    public void setWorkScheduleTitle(String workScheduleTitle) {
        this.workScheduleTitle = workScheduleTitle;
    }

    public String getWorkScheduleHours() {
        return workScheduleHours;
    }

    public void setWorkScheduleHours(String workScheduleHours) {
        this.workScheduleHours = workScheduleHours;
    }

    public String getJobOfferUrl() {
        return jobOfferUrl;
    }

    public void setJobOfferUrl(String jobOfferUrl) {
        this.jobOfferUrl = jobOfferUrl;
    }
}
