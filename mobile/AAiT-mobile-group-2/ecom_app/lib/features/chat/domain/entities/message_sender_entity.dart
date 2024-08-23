import 'package:equatable/equatable.dart';

class MessageSenderEntity extends Equatable{
  final String name;
  final String email;

  MessageSenderEntity({required this.name, required this.email});
  
  @override
  
  List<Object?> get props => [name, email];
}