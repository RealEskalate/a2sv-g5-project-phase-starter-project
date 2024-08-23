


import 'package:ecommers/core/network/check_connectivity.dart';
import 'package:ecommers/features/ecommerce/Data/data_source/local_data_source.dart';
import 'package:ecommers/features/ecommerce/Data/data_source/remote_data_source.dart';
import 'package:ecommers/features/ecommerce/Domain/repositories/ecommerce_repositories.dart';
import 'package:ecommers/features/ecommerce/Domain/usecase/ecommerce_usecase.dart';
import 'package:ecommers/features/ecommerce/presentation/state/image_input_display/image_bloc.dart';
import 'package:ecommers/features/ecommerce/presentation/state/input_button_activation/button_bloc.dart';
import 'package:ecommers/features/ecommerce/presentation/state/product_bloc/product_bloc.dart';
import 'package:ecommers/features/ecommerce/presentation/state/user_states/login_user_states_bloc.dart';
import 'package:ecommers/features/login/data/datasource/remote_datasource.dart';
import 'package:ecommers/features/login/data/repositories/login_repo_impl.dart';
import 'package:ecommers/features/login/domain/repositories/login_repositories.dart';
import 'package:ecommers/features/login/domain/usecase/login_usecase.dart';
import 'package:ecommers/features/login/presentation/state/Login_Registration/login_registration_bloc.dart';
import 'package:ecommers/features/login/presentation/state/login/login_bloc.dart';
import 'package:http/http.dart' as http;
import 'package:internet_connection_checker/internet_connection_checker.dart';
import 'package:mockito/annotations.dart';
import 'package:shared_preferences/shared_preferences.dart';

@GenerateMocks(

  [
    EcommerceRepositories,
    EcommerceRemoteDataSourceImpl,
    NetworkInfoImpl,
    InternetConnectionChecker,
    SharedPreferences,
    LocalDataSourceImpl,
    EcommerceUsecase,
    ProductBloc,
    ImageBloc,
    ButtonBloc,
    LoginRepositories,
    LoginRepoImpl,
    RemoteDatasourceImpl,
    LoginUserStatesBloc,
    LoginUseCase,
    LoginBloc,
    LoginRegistrationBloc
  ],
  customMocks :[MockSpec<http.Client>(as : #MockHttpClient)],

)
@GenerateMocks([http.Client])
void main() {}

