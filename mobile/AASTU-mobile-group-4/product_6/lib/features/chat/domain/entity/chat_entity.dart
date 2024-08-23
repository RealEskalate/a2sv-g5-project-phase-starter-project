import 'package:equatable/equatable.dart';
import 'message_entity.dart';
import '../../../auth/domain/entity/auth_entity.dart';

class ChatEntity extends Equatable {
  final String chatId;
  final UserEntity sender;
  final UserEntity receiver;

  ChatEntity({
    required this.chatId,
    required this.sender,
    required this.receiver,
   
  });

  @override
  List<Object?> get props => [chatId, sender, receiver];
}
