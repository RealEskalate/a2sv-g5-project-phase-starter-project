import '../../../authentication/domain/entities/user_data.dart';

class Chat {
  final String id;
  final UserEntity user;

  Chat({
    required this.id,
    required this.user,
  });
}