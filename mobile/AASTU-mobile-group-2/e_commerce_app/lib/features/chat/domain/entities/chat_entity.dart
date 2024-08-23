import 'package:e_commerce_app/features/auth/domain/entities/user.dart';
import 'package:equatable/equatable.dart';

class ChatEntity extends Equatable {
  final String id;
  final User sender;
  final User reciever;

  const ChatEntity({
    required this.id,
    required this.sender,
    required this.reciever,
   
  });
  @override
  List<Object?> get props => [id, sender, reciever];

}