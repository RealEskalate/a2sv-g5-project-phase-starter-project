

import '../../features/auth/data/datasource/auth_local_datasource/login_local_datasource.dart';
import '../../features/auth/data/datasource/remote_datasource/auth_remote_datasource.dart';
import '../../features/auth/data/repository_impl/auth_repo_impl.dart';
import '../../features/auth/domain/repository/auth_repo.dart';
import '../../features/auth/domain/usecase/log_out_usecase.dart';
import '../../features/auth/domain/usecase/login_usecase.dart';
import '../../features/auth/domain/usecase/signUp_usecase.dart';
import '../../features/auth/presentation/bloc/auth_bloc.dart';

import 'injection.dart';

class AuthInjection {
  init() {
    // Bloc
    sl.registerFactory<AuthBloc>(() => AuthBloc(sl(), sl(),sl()));

    // UseCase
    sl.registerLazySingleton<SignUpUseCase>(() => SignUpUseCase(sl()));
    sl.registerLazySingleton(() => LogInUseCase(sl()));
    sl.registerLazySingleton(() => LogoutUsecase(repository: sl()));

    // Repository
    sl.registerLazySingleton<AuthRepository>(
        () => AuthRepoImpl(sl(), sl(), sl()));

    // Data Source
    sl.registerLazySingleton<AuthRemoteDataSource>(
        () => AuthRemoteDataSourceImpl(client: sl()));
    sl.registerLazySingleton<UserLogInLocalDataSource>(
        () => UserLogInLocalDataSourceImpl(sharedPreferences: sl()));
  }
}
