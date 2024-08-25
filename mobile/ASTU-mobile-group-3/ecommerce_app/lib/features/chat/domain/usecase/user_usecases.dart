import '../../../auth/domain/entities/user_entity.dart';
import '../repository/chat_repository.dart';

class UserUsecases {
  final UserRepository userRepository;

  UserUsecases({required this.userRepository});
    Future<List<UserEntity>> callGetUsers()async{
      return await userRepository.getUsers();
    }


}