import 'package:flutter_bloc/flutter_bloc.dart';

import '../../features/auth/domain/entities/user_entity.dart';

class UserCubit extends Cubit<UserEntity?> {
  UserCubit() : super(null);

  void updateUser(UserEntity user) {
    emit(user);
  }
}
