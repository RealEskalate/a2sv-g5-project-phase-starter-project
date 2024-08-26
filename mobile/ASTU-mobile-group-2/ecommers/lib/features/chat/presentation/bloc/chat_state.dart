

import 'package:equatable/equatable.dart';

import '../../domain/entity/chat_entity.dart';

abstract class ChatState  extends Equatable{}

class ChatInitialState extends ChatState {
  ChatInitialState();
  
  @override

  List<Object?> get props => [];

  
}


class ChatMessageGetSuccess extends ChatState {
  final List<ChatEntity> chatEntity;

  ChatMessageGetSuccess ({
    required this.chatEntity
  });
  
  @override

  List<Object?> get props => [chatEntity];
  
}

class ChatDeleteSuccess extends ChatState {
  final bool chatDeleted;

  ChatDeleteSuccess ({
    required this.chatDeleted
  });


  @override

  List<Object?> get props => [chatDeleted];
}

class ChatErrorState extends ChatState {
  final String errorMessage;
  ChatErrorState ({
    required this.errorMessage
  });
  
  @override
  
  List<Object?> get props => [errorMessage];
  
}

class ChatLoadingState extends ChatState{
  ChatLoadingState();
  
  @override
  List<Object?> get props => [];

}

class ChatInitiatState extends ChatState {
  final bool chatInitiated;
  ChatInitiatState ({
    required this.chatInitiated
  });
  
  @override

  List<Object?> get props =>  [chatInitiated];

  
}

class ChatByIdSuccess extends ChatState {
  final ChatEntity chatEntity;

  ChatByIdSuccess ({
    required this.chatEntity
  });

  @override

  List<Object?> get props => [chatEntity];
}