import 'package:equatable/equatable.dart';

import '../../../authentication/domain/entity/user.dart';

class Chat extends Equatable{
  final String id;
  final User user1;
  final User user2;

  Chat({
    required this.id,
    required this.user1,  
    required this.user2,  
  });
  
  @override
  List<Object?> get props => [id, user1, user2];
}

