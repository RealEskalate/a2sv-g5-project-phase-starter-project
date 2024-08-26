import 'package:ecommerce_app/core/network/network_info.dart';
import 'package:ecommerce_app/core/validator/validator.dart';
import 'package:ecommerce_app/features/auth/data/data_source/auth_local_data_source.dart';
import 'package:ecommerce_app/features/auth/data/data_source/remote_auth_data_source.dart';
import 'package:ecommerce_app/features/auth/domain/repositories/auth_repository.dart';
import 'package:ecommerce_app/features/auth/domain/usecases/get_me_usecase.dart';
import 'package:ecommerce_app/features/auth/domain/usecases/log_in_usecase.dart';
import 'package:ecommerce_app/features/auth/domain/usecases/log_out_usecase.dart';
import 'package:ecommerce_app/features/auth/domain/usecases/sign_up_usecase.dart';
import 'package:ecommerce_app/features/auth/presentation/bloc/auth_bloc.dart';
import 'package:ecommerce_app/features/auth/presentation/bloc/cubit/user_input_validation_cubit.dart';
import 'package:ecommerce_app/features/chat/domain/repository/chat_repository.dart';
import 'package:ecommerce_app/features/chat/domain/usecases/AcknowledgeMessageDeliveryUseCase.dart';
import 'package:ecommerce_app/features/chat/domain/usecases/CreateChatRoomUseCase.dart';
import 'package:ecommerce_app/features/chat/domain/usecases/OnMessageReceivedUseCase.dart';
import 'package:ecommerce_app/features/chat/domain/usecases/RetrieveChatRoomsUseCase.dart';
import 'package:ecommerce_app/features/chat/domain/usecases/RetrieveMessagesUseCase.dart';
import 'package:ecommerce_app/features/chat/domain/usecases/SendMessageUseCase.dart';
import 'package:ecommerce_app/features/product/data/data_resources/local_product_data_source.dart';
import 'package:ecommerce_app/features/product/data/data_resources/remote_product_data_source.dart';
import 'package:ecommerce_app/features/product/domain/repositories/product_repository.dart';
import 'package:ecommerce_app/features/product/domain/usecases/delete_product_usecase.dart';
import 'package:ecommerce_app/features/product/domain/usecases/get_all_products_usecase.dart';
import 'package:ecommerce_app/features/product/domain/usecases/get_product_usecase.dart';
import 'package:ecommerce_app/features/product/domain/usecases/insert_product_usecase.dart';
import 'package:ecommerce_app/features/product/domain/usecases/update_product_usecase.dart';
import 'package:ecommerce_app/features/product/presentation/bloc/cubit/input_validation_cubit.dart';
import 'package:ecommerce_app/features/product/presentation/bloc/product_bloc.dart';
import 'package:http/http.dart' as http;
import 'package:internet_connection_checker/internet_connection_checker.dart';
import 'package:mockito/annotations.dart';
import 'package:shared_preferences/shared_preferences.dart';

@GenerateMocks(
  [
    ProductRepository,
    NetworkInfo,
    RemoteProductDataSource,
    LocalProductDataSource,
    InternetConnectionChecker,
    SharedPreferences,
    GetAllProductUseCase,
    GetProductUseCase,
    UpdateProductUsecase,
    InsertProductUseCase,
    DeleteProductUseCase,
    AuthRepository,
    RemoteAuthDataSource,
    AuthLocalDataSource,
    LogInUsecase,
    SignUpUsecase,
    LogOutUsecase,
    InputDataValidator,
    AuthBloc,
    ProductBloc,
    InputValidationCubit,
    GetMeUsecase,
    UserInputValidationCubit,
    ChatRepository,
    SendMessageUseCase,
    RetrieveMessagesUseCase,
    RetrieveChatRoomsUseCase,
    OnMessageReceivedUseCase,
    CreateChatRoomUseCase,
    AcknowledgeMessageDeliveryUseCase
  ],
  customMocks: [MockSpec<http.Client>(as: #MockHttpClient)],
)
void main() {}