package com.xws.agentska.mapper;

import com.xws.agentska.dto.UserExperienceDto;
import com.xws.agentska.model.UserExperience;
import org.modelmapper.ModelMapper;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component("userExperienceModelMapper")
public class UserExperienceModelMapper<E extends UserExperience,T> {
    @Autowired
    private ModelMapper modelMapper;

    public E convertToEntity(UserExperienceDto dto, Class<E> entityClass) {
        return modelMapper.map(dto,entityClass);
    }
}
