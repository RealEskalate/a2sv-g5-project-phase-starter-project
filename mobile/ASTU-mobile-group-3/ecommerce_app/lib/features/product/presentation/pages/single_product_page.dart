import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../core/themes/themes.dart';
import '../../../chat/presentation/bloc/chat_bloc.dart';
import '../../../chat/presentation/bloc/chat_event.dart';
import '../../../chat/presentation/bloc/chat_state.dart';
import '../bloc/product_bloc.dart';
import '../bloc/product_events.dart';
import '../bloc/product_states.dart';
import '../widgets/loading_dialog.dart';
import '../widgets/product_widgets.dart';
import 'update_product_page.dart';

// ignore: must_be_immutable
class SingleProduct extends StatelessWidget {
  static String routes = '/single_product';
  String? sellerId;
  List<List<int>> size = [
    [1, 39],
    [0, 40],
    [0, 41],
    [0, 42],
    [0, 43],
    [0, 44],
    [0, -1]
  ];
  SingleProduct({super.key});
  @override
  Widget build(BuildContext context) {
    final mediaQuery = MediaQuery.of(context);
    const textStyle = TextStyle(
      fontFamily: 'poppins',
      fontWeight: FontWeight.w500,
      fontSize: 16,
    );
    return Scaffold(
      body: SingleChildScrollView(
        child: BlocListener<ProductBloc, ProductStates>(
          listener: (context, state) {
            if (state is SuccessfullState) {
              BlocProvider.of<ProductBloc>(context).add(LoadAllProductEvents());
              Navigator.pop(context);
            } else if (state is ErrorState) {
              ScaffoldMessenger.of(context).showSnackBar(
                SnackBar(
                  content: Text(
                    state.message,
                  ),
                ),
              );
            }
          },
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.stretch,
            children: [
              Stack(
                alignment: Alignment.topLeft,
                children: [
                  BlocBuilder<ProductBloc, ProductStates>(
                    builder: (context, state) {
                      if (state is LoadedSingleProductState) {
                        return Container(
                          height: 250,
                          width: double.infinity,
                          decoration: const BoxDecoration(
                            color: MyTheme.ecGrey,
                          ),
                          child: Image.network(
                            fit: BoxFit.fill,
                            state.productEntity.imageUrl,
                            loadingBuilder: (context, myChild, progress) {
                              if (progress == null) {
                                return myChild;
                              } else {
                                return const SizedBox();
                              }
                            },
                            errorBuilder: (context, exception, trace) {
                              return const SizedBox();
                            },
                          ),
                        );
                      } else if (state is ErrorState) {
                        return Container(
                          decoration:
                              const BoxDecoration(color: MyTheme.ecGrey),
                          height: 250,
                          width: double.infinity,
                          child: Center(
                            child: Text(
                              state.message,
                              style: const TextStyle(
                                color: MyTheme.ecRed,
                              ),
                            ),
                          ),
                        );
                      } else {
                        return Container(
                          height: 250,
                          width: double.infinity,
                          decoration:
                              const BoxDecoration(color: MyTheme.ecGrey),
                          child: const Center(
                            child: CircularProgressIndicator(),
                          ),
                        );
                      }
                    },
                  ),
                  SafeArea(
                    child: Padding(
                      padding: const EdgeInsets.all(10),
                      child: IconButton(
                        onPressed: () {
                          BlocProvider.of<ProductBloc>(context)
                              .add(LoadAllProductEvents());
                          Navigator.pop(context);
                        },
                        icon: const Icon(
                          Icons.chevron_left,
                          color: Colors.blue,
                        ),
                        style:
                            IconButton.styleFrom(backgroundColor: Colors.white),
                      ),
                    ),
                  ),
                ],
              ),
              Padding(
                padding: const EdgeInsets.symmetric(horizontal: 20, vertical: 10),
                child: Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    BlocBuilder<ProductBloc, ProductStates>(
                      builder: (context, state) {
                        if (state is LoadedSingleProductState) {
                          sellerId = state.productEntity.seller.id;
                          return ConstrainedBox(
                            constraints: BoxConstraints(
                                maxWidth:
                                    MediaQuery.of(context).size.width / 2),
                            child: Text(
                              state.productEntity.description,
                              style: const TextStyle(
                                  color: Colors.grey,
                                  fontSize: 18,
                                  fontFamily: 'poppins',
                                  overflow: TextOverflow.ellipsis),
                            ),
                          );
                        } else {
                          return const Text('...');
                        }
                      },
                    ),
                    const Wrap(
                      children: [
                        Icon(
                          Icons.star,
                          color: Color.fromARGB(255, 255, 184, 100),
                        ),
                        Text(
                          '(5.0)',
                          style: TextStyle(
                            color: Colors.grey,
                            fontSize: 18,
                            fontFamily: 'poppins',
                          ),
                        ),
                      ],
                    ),
                  ],
                ),
              ),
              Padding(
                padding:
                    const EdgeInsets.symmetric(horizontal: 20, vertical: 2),
                child: BlocBuilder<ProductBloc, ProductStates>(
                  builder: (context, state) {
                    if (state is LoadedSingleProductState) {
                      return Row(
                        mainAxisAlignment: MainAxisAlignment.spaceBetween,
                        children: [
                          Text(
                            state.productEntity.name,
                            style: const TextStyle(
                                fontFamily: 'poppins',
                                fontSize: 24,
                                fontWeight: FontWeight.w500),
                          ),
                          Text(
                            '${state.productEntity.price}\$',
                            style: textStyle,
                          ),
                        ],
                      );
                    } else {
                      return const Row(
                        mainAxisAlignment: MainAxisAlignment.spaceBetween,
                        children: [
                          Text(
                            '...',
                            style: TextStyle(
                                fontFamily: 'poppins',
                                fontSize: 24,
                                fontWeight: FontWeight.w500),
                          ),
                          Text(
                            '0\$',
                            style: textStyle,
                          ),
                        ],
                      );
                    }
                  },
                ),
              ),
              const Padding(
                padding: EdgeInsets.symmetric(
                  horizontal: 20,
                  vertical: 5,
                ),
                child: Text(
                  'Size:',
                  style: TextStyle(
                    fontFamily: 'poppins',
                    fontWeight: FontWeight.w500,
                    fontSize: 20,
                    color: Color(0xFF3E3E3E),
                  ),
                ),
              ),
              SizedBox(
                width: mediaQuery.size.width,
                height: 70,
                child: Padding(
                  padding: const EdgeInsets.symmetric(horizontal: 20),
                  child: ListView.builder(
                    scrollDirection: Axis.horizontal,
                    itemCount: size.length - 1,
                    itemBuilder: (context, index) {
                      return GestureDetector(
                        onTap: () {},
                        child: Container(
                            margin: EdgeInsets.fromLTRB(5, 4, 5, 4),
                            height: 60,
                            width: 60,
                           
                            decoration: BoxDecoration(
                                color: (size[index][0] == 0)
                                    ? const Color(0XFFFFFFFF)
                                    : const Color(0XFF3F51F3),
                                boxShadow:  [
                                  BoxShadow(
                                      color: Colors.black.withOpacity(0.1),
                                      spreadRadius: 0.5,
                                      blurRadius: 1,
                                      offset: Offset(0,2))
                                ],
                                borderRadius: BorderRadius.circular(10)),
                            child: Center(
                                child: Text(
                              size[index][1].toString(),
                              style: TextStyle(
                                  color: (size[index][0] == 0)
                                      ? Colors.black
                                      : Colors.white,
                                  fontWeight: FontWeight.w500,
                                  fontSize: 19),
                            ))),
                      );
                    },
                  ),
                ),
              ),
              Padding(
                padding: const EdgeInsets.symmetric(
                  horizontal: 20,
                  vertical: 10,
                ),
                child: BlocBuilder<ProductBloc, ProductStates>(
                  builder: (context, state) {
                    if (state is LoadedSingleProductState) {
                      return Text(
                        state.productEntity.description,
                        style: const TextStyle(
                          fontFamily: 'poppins',
                          fontSize: 12,
                        ),
                      );
                    } else {
                      return const LinearProgressIndicator();
                    }
                  },
                ),
              ),
              Padding(
                padding:
                    const EdgeInsets.symmetric(horizontal: 10, vertical: 10),
                child: Row(
                  mainAxisAlignment: MainAxisAlignment.spaceAround,
                  children: [
                    Expanded(
                      child: OutlineCustomButton(
                        press: () {
                          final result =
                              BlocProvider.of<ProductBloc>(context).state;
                          if (result is LoadedSingleProductState) {
                            String id = result.productEntity.id;
                            BlocProvider.of<ProductBloc>(context)
                                .add(DeleteProductEvent(id: id));
                          }
                        },
                        label: 'DELETE',
                      ),
                    ),
                    SizedBox(
                      width: 10,
                    ),
                    Expanded(
                      child: BlocConsumer<ChatBloc, ChatState>(
                        listener: (context, state) {
                          if (state is ChatLoaded) {
                            Navigator.pop(context);

                            ScaffoldMessenger.of(context).showSnackBar(
                                const SnackBar(
                                    content: Text('Chat Room Created')));
                          } else if (state is ChatLoading) {
                            showDialog(
                                context: context,
                                builder: (_) {
                                  return const LoadingDialog();
                                });
                          }
                        },
                        builder: (context, state) {
                          return FillCustomButton(
                              press: () {
                                if (sellerId != null) {
                                  final chatBloc =
                                      BlocProvider.of<ChatBloc>(context);

                                  chatBloc.add(CreateChatRoom(sellerId!));
                                }
                              },
                              label: 'CHAT');
                        },
                      ),
                    ),
                    SizedBox(
                      width: 10,
                    ),
                    Expanded(
                      child: FillCustomButton(
                          press: () {
                            final state =
                                BlocProvider.of<ProductBloc>(context).state;
                            if (state is LoadedSingleProductState) {
                              BlocProvider.of<ProductBloc>(context).add(
                                GetSingleProductEvents(
                                    id: state.productEntity.id),
                              );
                            }

                            Navigator.pushNamed(
                                context, UpdateProductPage.routes);
                          },
                          label: 'UPDATE'),
                    ),
                  ],
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
