package com.xws.agentska.dto;

public class SalaryExperienceDto extends UserExperienceDto {
    private double salary;

    public SalaryExperienceDto() {
    }

    public SalaryExperienceDto(double salary) {
        this.salary = salary;
    }

    public double getSalary() {
        return salary;
    }

    public void setSalary(double salary) {
        this.salary = salary;
    }
}
