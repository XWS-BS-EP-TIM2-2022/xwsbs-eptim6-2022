package com.xws.agentska.model;

import com.xws.agentska.model.enumerations.DifficultyLevels;

import javax.persistence.Entity;
import javax.persistence.EnumType;
import javax.persistence.Enumerated;

@Entity

public class InterviewExperience extends Comment{
    @Enumerated(EnumType.STRING)
    private DifficultyLevels difficultyLevel;
    private double selectionProcessDuration;

    public InterviewExperience() {
    }

    public DifficultyLevels getDifficultyLevel() {
        return difficultyLevel;
    }

    public void setDifficultyLevel(DifficultyLevels difficultyLevel) {
        this.difficultyLevel = difficultyLevel;
    }

    public double getSelectionProcessDuration() {
        return selectionProcessDuration;
    }

    public void setSelectionProcessDuration(double selectionProcessDuration) {
        this.selectionProcessDuration = selectionProcessDuration;
    }
}
