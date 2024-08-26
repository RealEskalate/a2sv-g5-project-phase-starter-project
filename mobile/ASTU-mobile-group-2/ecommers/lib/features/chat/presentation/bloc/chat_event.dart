




import 'package:equatable/equatable.dart';

abstract class ChatEvent  extends Equatable{}


class OnGetAllChat extends ChatEvent{

  @override
  List<Object?> get props => [];
  
}


class OnDeleteChat extends ChatEvent {
  final String chatId;
  OnDeleteChat ({
    required this.chatId
  });

  @override
  List<Object?> get props => [];

  
}

class OnInitiatChat extends ChatEvent {
  final String userId;
  

  OnInitiatChat ({
    required this.userId
  });

  @override

  List<Object?> get props => [];
  
}

class OnGetChatById extends ChatEvent {
  final String chatId;

  OnGetChatById ({
    required this.chatId
  }); 


  @override
  List<Object?> get props => [];
  }