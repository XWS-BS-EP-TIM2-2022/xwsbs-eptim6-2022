package com.xws.agentska.mapper;

import com.xws.agentska.dto.CommentDto;
import com.xws.agentska.model.Comment;
import org.springframework.stereotype.Component;

@Component("commentModelMapper")
public class CommentModelMapper extends CustomModelMapperAbstract<Comment, CommentDto>{
    @Override
    public CommentDto convertToDto(Comment entity) {
        return null;
    }

    @Override
    public Comment convertToEntity(CommentDto dto) {
        return modelMapper.map(dto,Comment.class);
    }
}
