package com.xws.agentska.mapper;

import com.xws.agentska.dto.UserDto;
import com.xws.agentska.model.User;

public class UsersModelMapper extends CustomModelMapperAbstract<User, UserDto> {
    @Override
    public UserDto convertToDto(User entity) {
        return modelMapper.map(entity,UserDto.class);
    }

    @Override
    public User convertToEntity(UserDto dto) {
        return modelMapper.map(dto,User.class);
    }
}
