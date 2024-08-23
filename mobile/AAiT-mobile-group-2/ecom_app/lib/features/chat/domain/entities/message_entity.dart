import 'package:ecom_app/features/chat/domain/entities/message_sender_entity.dart';
import 'package:equatable/equatable.dart';

class MessageEntity extends Equatable{
  final MessageSenderEntity messageSenderEntity;
  final String content;

  MessageEntity({required this.messageSenderEntity, required this.content});
  
  @override
  List<Object?> get props => [messageSenderEntity, content];
}