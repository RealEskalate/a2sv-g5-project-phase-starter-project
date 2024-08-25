import 'package:bloc_test/bloc_test.dart';
import 'package:ecommerce_app/features/auth/presentation/bloc/auth_bloc.dart';
import 'package:ecommerce_app/features/auth/presentation/bloc/cubit/user_input_validation_cubit.dart';
import 'package:ecommerce_app/features/product/presentation/bloc/cubit/input_validation_cubit.dart';
import 'package:ecommerce_app/features/product/presentation/bloc/product_bloc.dart';
import 'package:ecommerce_app/features/product/presentation/bloc/product_events.dart';
import 'package:ecommerce_app/features/product/presentation/bloc/product_states.dart';

class MockAuthBloc extends MockBloc<AuthEvent, AuthState> implements AuthBloc {}

class MockUserInput extends MockCubit<UserInputValidationState>
    implements UserInputValidationCubit {}

class MockInputValidator extends MockCubit<InputValidationState>
    implements InputValidationCubit {}

class MockProductBloc extends MockBloc<ProductEvents, ProductStates>
    implements ProductBloc {}
