package com.xws.agentska.dto;

import com.xws.agentska.model.enumerations.DifficultyLevels;

public class InterviewExperienceDto extends CommentDto {
    private DifficultyLevels difficultyLevel;
    private double selectionProcessDuration;

    public InterviewExperienceDto() {
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
