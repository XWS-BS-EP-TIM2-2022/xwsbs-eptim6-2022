package com.xws.agentska.model;

import javax.persistence.Entity;

@Entity
public class SalaryExperience extends UserExperience {
    private double salary;

    public SalaryExperience() {
    }

    public double getSalary() {
        return salary;
    }

    public void setSalary(double salary) {
        this.salary = salary;
    }
}
