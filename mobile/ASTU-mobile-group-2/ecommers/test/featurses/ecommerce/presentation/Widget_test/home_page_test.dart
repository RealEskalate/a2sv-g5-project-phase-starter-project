import 'dart:io';

import 'package:bloc_test/bloc_test.dart';
import 'package:dartz/dartz.dart';
import 'package:ecommers/core/utility/socket_impl.dart';
import 'package:ecommers/features/chat/domain/entity/chat_entity.dart';
import 'package:ecommers/features/chat/domain/entity/message_entity.dart';
import 'package:ecommers/features/chat/domain/usecase/chat_usecase.dart';
import 'package:ecommers/features/chat/presentation/bloc/chat_bloc.dart';
import 'package:ecommers/features/chat/presentation/bloc/chat_event.dart';
import 'package:ecommers/features/chat/presentation/bloc/chat_state.dart';
import 'package:ecommers/features/ecommerce/Domain/entity/ecommerce_entity.dart';
import 'package:ecommers/features/ecommerce/presentation/UI/home/home.dart';
import 'package:ecommers/features/ecommerce/presentation/UI/home/product_image.dart';
import 'package:ecommers/features/ecommerce/presentation/state/product_bloc/product_bloc.dart';
import 'package:ecommers/features/ecommerce/presentation/state/product_bloc/product_event.dart';
import 'package:ecommers/features/ecommerce/presentation/state/product_bloc/product_state.dart';
import 'package:ecommers/features/ecommerce/presentation/state/user_states/login_user_states_bloc.dart';
import 'package:ecommers/features/ecommerce/presentation/state/user_states/login_user_states_event.dart';
import 'package:ecommers/features/ecommerce/presentation/state/user_states/login_user_states_state.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:get_it/get_it.dart';
import 'package:mocktail/mocktail.dart';

import '../../../../helper/test_hlper.mocks.dart';
class MockProductBloc extends MockBloc<ProductEvent, ProductState> implements ProductBloc {}
class MockLoginUserStatesBloc extends MockBloc<LoginUserStatesEvent, LoginUserStates> implements LoginUserStatesBloc {}
class MockSocketService extends Mock implements SocketService {}
class MockChatBloc extends MockBloc<ChatEvent, ChatState> implements ChatBloc {}
class MockChatUsecase extends Mock implements ChatUsecase {}
void main() {
  late MockSocketService mockSocketService;
  late MockProductBloc mockProductBloc;
  late MockLoginUserStatesBloc mockLoginUserStatesBloc;
  late MockChatUsecase mockChatUsecase;
  late MockChatBloc mockChatBloc;
  final MessageEntity messageEntity = MessageEntity (
      messages: [{'id' : '', 'content': ''}], messageId: ''
    );
     final List<ChatEntity> chatEntity = [
      ChatEntity(
        senderName: '', recieverId: '', senderId: '', recieverName: '', chatId: '', messages: messageEntity
      ),
      
    ];

  setUp(() {
    mockSocketService = MockSocketService();
    mockChatUsecase = MockChatUsecase();
    mockProductBloc = MockProductBloc();
    mockChatBloc = MockChatBloc();
    mockLoginUserStatesBloc = MockLoginUserStatesBloc();
    
    // Stubbing connect method to return a completed Future
    when(() => mockSocketService.connect()).thenAnswer((_) async => Future.value());
    
    GetIt.instance.registerSingleton<SocketService>(mockSocketService);
    HttpOverrides.global = null;
  });

  tearDown(() {
    mockProductBloc.close();
    mockLoginUserStatesBloc.close();
    GetIt.instance.reset();
  });

  testWidgets('HomeScreen displays products correctly', (WidgetTester tester) async {
    // Arrange
    when(() => mockChatUsecase.getMychats()).thenAnswer(
      (_) async => Right(chatEntity) // Make sure this matches the expected return type
    );

    whenListen(
      mockProductBloc,
      Stream.fromIterable([
        LoadingState(),
        const LoadedAllProductState(
          products: [
            EcommerceEntity(id: '1', name: 'Product 1', description: 'Description 1', imageUrl: 'https://www.simplilearn.com/ice9/free_resources_article_thumb/what_is_image_Processing.jpg', price: 100, sellerId: '', sellerName: ''),
            EcommerceEntity(id: '2', name: 'Product 2', description: 'Description 2', imageUrl: 'https://www.simplilearn.com/ice9/free_resources_article_thumb/what_is_image_Processing.jpg', price: 200, sellerId: '', sellerName: ''),
          ],
        ),
      ]),
      initialState: ProductIntialState(),
    );

    whenListen(
      mockLoginUserStatesBloc,
      Stream.fromIterable([
        LeftUserStates(), // or any other initial state
      ]),
      initialState: LeftUserStates(),
    );

    whenListen(
      mockChatBloc,
      Stream.fromIterable([
        ChatMessageGetSuccess(chatEntity: chatEntity), // or any other initial state
      ]),
      initialState: ChatInitialState(),
    );

    // Act
    await tester.pumpWidget(
      MaterialApp(
        home: MultiBlocProvider(
          providers: [
            BlocProvider<ChatBloc>(
              create: (context) => mockChatBloc,
            ),
            BlocProvider<ProductBloc>.value(value: mockProductBloc),
            BlocProvider<LoginUserStatesBloc>.value(value: mockLoginUserStatesBloc),
          ],
          child: const HomeScreen(),
        ),
      ),
    );

    // Process the second state (LoadedAllProductState)
    await tester.pumpAndSettle();
    // Assert
    expect(find.text('Available Products'), findsOneWidget);
  });

  testWidgets('HomeScreen displays error state', (WidgetTester tester) async {
    // Arrange
    when(() => mockProductBloc.state).thenAnswer(
      (_) => const ProductErrorState(messages: 'try again')
    );

    when(() => mockLoginUserStatesBloc.state).thenAnswer(
      (_) => LeftUserStates() // or any other initial state
    );

    // Act
    await tester.pumpWidget(
      MaterialApp(
        home: MultiBlocProvider(
          providers: [
            BlocProvider<ChatBloc>(
              create: (context) => mockChatBloc,
            ),
            BlocProvider<ProductBloc>.value(value: mockProductBloc),
            BlocProvider<LoginUserStatesBloc>.value(value: mockLoginUserStatesBloc),
          ],
          child: const HomeScreen(),
        ),
      ),
    );

    // Process the second state (ProductErrorState)
    await tester.pumpAndSettle();

    
    expect(find.byType(ElevatedButton), findsOneWidget);
    
  });

  testWidgets('HomeScreen must show loading state', (WidgetTester tester) async {
    // Arrange
    when(() => mockProductBloc.state).thenAnswer(
      (_) => LoadingState()
    );

    when(() => mockLoginUserStatesBloc.state).thenAnswer(
      (_) => LeftUserStates() // or any other initial state
    );

    whenListen(
      mockChatBloc,
      Stream.fromIterable([
        ChatMessageGetSuccess(chatEntity: chatEntity), // or any other initial state
      ]),
      initialState: ChatInitialState(),
    );

    // Act
    await tester.pumpWidget(
      MaterialApp(
        home: MultiBlocProvider(
          providers: [
            BlocProvider<ChatBloc>(
              create: (context) => mockChatBloc,
            ),
            BlocProvider<ProductBloc>.value(value: mockProductBloc),
            BlocProvider<LoginUserStatesBloc>.value(value: mockLoginUserStatesBloc),
          ],
          child: const HomeScreen(),
        ),
      ),
    );

    // Process the state (LoadingState)
    await tester.pumpAndSettle();

    expect(find.byKey(const Key('loading')), findsOneWidget);
  });
}