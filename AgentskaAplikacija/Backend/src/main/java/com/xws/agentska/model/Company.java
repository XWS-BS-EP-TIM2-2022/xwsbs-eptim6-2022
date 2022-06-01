package com.xws.agentska.model;

import com.xws.agentska.model.enumerations.Status;

import javax.persistence.*;
import java.util.HashSet;
import java.util.Set;

@Entity
public class Company {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private String name;
    private String description;
    private String culture;
    @Embedded
    private ContactInfo contactInfo;
    private String web;
    private int yearOfEstablishment;
    private Status status;
    @OneToOne
    @JoinColumn(name = "owner_id")
    private User owner;
    @OneToMany(fetch = FetchType.LAZY, cascade = CascadeType.ALL)
    private Set<JobOffer> jobOffers;
    @OneToOne(cascade = CascadeType.ALL)
    private ApiConnection apiConnection;

    public User getOwner() {
        return owner;
    }

    public Company() {
    }

    public ContactInfo getContactInfo() {
        return contactInfo;
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public String getCulture() {
        return culture;
    }

    public void setCulture(String culture) {
        this.culture = culture;
    }

    public void setContactInfo(ContactInfo contactInfo) {
        this.contactInfo = contactInfo;
    }

    public String getWeb() {
        return web;
    }

    public void setWeb(String web) {
        this.web = web;
    }

    public int getYearOfEstablishment() {
        return yearOfEstablishment;
    }

    public void setYearOfEstablishment(int yearOfEstablishment) {
        this.yearOfEstablishment = yearOfEstablishment;
    }

    public Status getStatus() {
        return status;
    }

    public void setStatus(Status status) {
        this.status = status;
    }

    public void setOwner(User owner) {
        this.owner = owner;
    }

    public Set<JobOffer> getJobOffers() {
        return jobOffers;
    }

    public void setJobOffers(Set<JobOffer> jobOffers) {
        this.jobOffers = jobOffers;
    }

    public ApiConnection getApiConnection() {
        return apiConnection;
    }

    public void setApiConnection(ApiConnection apiConnection) {
        this.apiConnection = apiConnection;
    }

    public void addNewJobOffer(JobOffer offer) {
        if (this.jobOffers == null) this.jobOffers = new HashSet<>();
        this.jobOffers.add(offer);
    }
}
