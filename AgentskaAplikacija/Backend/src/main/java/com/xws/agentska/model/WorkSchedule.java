package com.xws.agentska.model;

import javax.persistence.Embeddable;

@Embeddable
public class WorkSchedule {
    private String name;
    private int hoursPerWeek;

    public WorkSchedule() {
    }

    public WorkSchedule(String name, int hoursPerWeek) {
        this.name = name;
        this.hoursPerWeek = hoursPerWeek;
    }

}
