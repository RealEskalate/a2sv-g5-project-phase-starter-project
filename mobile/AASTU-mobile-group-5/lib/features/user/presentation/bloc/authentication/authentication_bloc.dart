import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';

import '../../../data/data_sources/user_local_data_source.dart';

part 'authentication_event.dart';
part 'authentication_state.dart';

class AuthenticationBloc extends Bloc<AuthenticationEvent, AuthenticationState> {
  final UserLocalDataSource localDataSource;

  AuthenticationBloc(this.localDataSource) : super(AuthenticationInitial()) {
    on<CheckCurrentStatus>((event, emit) async {
      final token = await localDataSource.getAccessToken();
      if (token != null) {
        emit(LoggedInState());
      } else {
        emit(LoggedOutState());
      }
    });

    on<LoggedOut>((event, emit) async {
      await localDataSource.deleteAccessToken();
      emit(LoggedOutState());
    });
  }
}