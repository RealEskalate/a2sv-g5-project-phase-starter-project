import 'package:equatable/equatable.dart';

class ChatEntity extends Equatable{
  final String id;
  final String user1;
  final String user2;
  

  ChatEntity({required this.id, required this.user1, required this.user2,
  });
  
  @override
  // TODO: implement props
  List<Object?> get props => [id, user1, user2,];


  Map<String, String> toJson() {
    return {
      '_id': id,
      'user1': user1,
      'user2': user2
    };

  }

  factory ChatEntity.fromJson(Map<String, dynamic> json) {
    return ChatEntity(
      id: json['id'],
      user1: json['user1'],
      user2: json['user2']
      );
  }
  
}


