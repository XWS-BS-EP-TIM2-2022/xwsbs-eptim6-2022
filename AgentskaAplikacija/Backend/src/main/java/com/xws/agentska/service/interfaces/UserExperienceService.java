package com.xws.agentska.service.interfaces;

import com.xws.agentska.model.Comment;
import com.xws.agentska.model.UserExperience;

import java.util.List;

public interface UserExperienceService extends Service<UserExperience,Long> {
    public List<Comment> findAllCommentsByCompanyId(long id);
}
